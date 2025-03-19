# frozen_string_literal: true

require "bundler/gem_tasks"
require "rspec/core/rake_task"

RSpec::Core::RakeTask.new(:spec)

require "rake/extensiontask"

task build: :compile

GEMSPEC = Gem::Specification.load("go_gem_yield_within_goroutine.gemspec")

Rake::ExtensionTask.new("go_gem_yield_within_goroutine", GEMSPEC) do |ext|
  ext.lib_dir = "lib/go_gem_yield_within_goroutine"
end

task default: %i[clobber compile spec]
